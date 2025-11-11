import Cards from './Cards';
import { useDroppable } from '@dnd-kit/core';
import './Columns.css';

const Columns = ({title, tasks, status, onEdit, onDelete}) =>{
    
    //controle drop
    const filteredTasks = tasks.filter(task => task.status === status);
    const { setNodeRef } = useDroppable({
        id: status,
    })
    
    return(
        <div ref={setNodeRef} className="column">
            <h2>{title}</h2>
            <div className="columnCards">
                {filteredTasks.map((task) => {
                    return <Cards 
                    key={task.id} 
                    task={task}
                    onEdit={onEdit} //passa função edit pro card
                    onDelete={onDelete} /> //passa função delete pro card
                })}
            </div>
        </div>
    )
} 
export default Columns;