import {useState} from "react"
import './Editor.css'

const Editor =({task, onSave, onClose, mode='edit'}) =>{

    // controle do modo do componente
    const [EditCard, setEditCard] = useState(
        mode === 'add'
        ? {titulo: '', descricao: '', status: 'todo'}
        :{
            titulo: task.titulo,
            descricao: task.descricao,
            status: task.status
        }
    );

    // salva 
    const handleSave = () => {
        const updated = mode === 'edit'
            ? { ...EditCard, id: task.id } 
            : EditCard;

        onSave(updated);
    };

    // titulo dinâmico do componente
    const header = mode === 'add' ? 'Adicionar Tarefa:' : 'Editar Tarefa';

    return (
        <div className="overlay">
            <div className="content">
                <div className="contentHeader">
                    <h2>{header}</h2>
                    <input type="text"
                    id="cardTitle"
                    value={EditCard.titulo}
                    onChange={(e) => setEditCard({...EditCard, titulo: e.target.value})} 
                    placeholder="Nome da tarefa..."
                    maxLength={20}>
                    </input>   
                </div>
                <div className="contentMidle">
                    <h2>Descrição:</h2>
                    <input type="text"
                    id="cardDesc"
                    value={EditCard.descricao}
                    onChange={(e) => setEditCard({...EditCard, descricao: e.target.value})}
                    placeholder="Descrição da tarefa"
                    maxLength={50}>
                    </input>
                </div>
                <div className="contentFooter">
                    <button onClick={handleSave}>
                        {mode === 'add' ? 'Adicionar' : 'Salvar'}
                    </button>
                    <button onClick={onClose}>Cancelar</button>
                </div>
            </div>
        </div>
    )
}

export default Editor;