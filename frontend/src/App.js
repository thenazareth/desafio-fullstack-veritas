import {useState, useEffect} from 'react';
import {DndContext} from '@dnd-kit/core';
import Columns from './Columns';
import Editor from './Editor';
import './App.css';

function App() {

   const [tasks, setTasks] = useState([]);
   const API_URL = 'http://localhost:8080/tasks';

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      const response = await fetch(API_URL);
      const data = await response.json();
      setTasks(data);
    } catch (error) {
      console.error('Erro ao carregar tasks:', error);
    }
  };

  const ColumnsType = [
    { id: 'todo', title: 'A FAZER' },
    { id: 'inprogress', title: 'EM PROGRESSO' },
    { id: 'done', title: 'CONCLUÍDO' },
  ];

  

  const [modalState, setModalState] = useState({
    isOpen: false,
    mode: 'edit',
    task: null
  });

  const handleAddCard = () => {
    setModalState({
      isOpen: true,
      mode: 'add',
      task: null
    });
  };

  const handleEditCard = (task) => {
    setModalState({
      isOpen: true,
      mode: 'edit',
      task: task
    })
  }

  const handleSave = async (updatedTask) => {
    try {
      if (modalState.mode === 'edit') {
        // edit usa PUT
        const response = await fetch(`${API_URL}/${updatedTask.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(updatedTask),
        });
        
        if (response.ok) {
          const savedTask = await response.json();
          setTasks(prev => prev.map(task => 
            task.id === savedTask.id ? savedTask : task
          ));
          handleClose();
        }
      } else {
        // Add usa POST  
        const response = await fetch(API_URL, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(updatedTask),
        });
        
        if (response.ok) {
          const newTask = await response.json();
          setTasks(prev => [...prev, newTask]);
          handleClose();
        }
      }
    } catch (error) {
      console.error('Erro ao salvar task:', error);
    }
  };

  const handleClose = () => {
    setModalState({isOpen: false, mode:'edit', task:null});
  };

  const handleDeleteCard = async (id) => {
    try {
      const response = await fetch(`${API_URL}/${id}`, {
        method: 'DELETE',
      });
      
      if (response.ok) {
        setTasks(prev => prev.filter(task => task.id !== id));
      }
    } catch (error) {
      console.error('Erro ao deletar task:', error);
    }
  };

  function handleDragEnd(event) {
  const { active, over} = event;

  if (!over) return;

  const taskId = active.id;
  const newStatus = over.id;
  
  const taskToUpdate = tasks.find(task => task.id === taskId);
  if (!taskToUpdate) return;

  setTasks(prevTasks =>
    prevTasks.map(task =>
      task.id === taskId ? { ...task, status: newStatus } : task
    )
  );

  fetch(`${API_URL}/${taskId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      ...taskToUpdate,
      status: newStatus
    }),
  }).catch(error => {
    console.error('Erro ao atualizar no backend:', error);
    setTasks(prevTasks =>
      prevTasks.map(task =>
        task.id === taskId ? { ...task, status: taskToUpdate.status } : task
      )
    );
  });
}

  return (
    <div className="App">
      <div className="kanbanHeader">
        <div className="addSect">
          <h2>Kanban</h2>
          <button onClick={handleAddCard}>
            Adicionar tarefa
          </button>
        </div>
      </div>
      <div className="Kanban">
        <DndContext onDragEnd={handleDragEnd}>
          {ColumnsType.map((col) => (
            <Columns 
              key={col.id}
              title={col.title}
              tasks={tasks}
              status={col.id}
              onEdit={handleEditCard}
              onDelete={handleDeleteCard}
            />
          ))}
        </DndContext>
      </div>
      {modalState.isOpen && (
        <Editor
        task={modalState.task}
        onSave={handleSave}
        onClose={handleClose}
        mode={modalState.mode}/>
      )}
    </div>
  );
}

export default App;