import './Cards.css';
import EditIcon from './Edit.svg';
import Close from './close_icon.svg';
import { useDraggable } from '@dnd-kit/core';

const Cards = ({task, onEdit, onDelete}) =>{

    //controle drag
    const {attributes, listeners, setNodeRef, transform} = useDraggable({
        id: task.id,
    });

    //movimento do card
    const style = transform ? {
        transform: `translate(${transform.x}px, ${transform.y}px)`,
        cursor: 'grabbing',
    } : undefined;


    return(
        <div className="taskCard" 
        ref={setNodeRef}
        style={{ cursor: 'grab', ...style }} 
        {...attributes} 
        {...listeners}>
            <div className="taskHeader">
                <button id="editButton"
                onPointerDown={(e) => e.stopPropagation()} //para drag para clicar no botão
                onClick={() => onEdit(task)}>
                    <img src={EditIcon} alt="Edit" />
                </button>
                <button id="closeButton"
                onPointerDown={(e) => e.stopPropagation()} //para drag para clicar no botão
                onClick={() => onDelete(task.id)}>
                    <img src={Close} alt="Close" />
                </button>
            </div>
            <h2>{task.titulo}</h2>
            <p>{task.descricao}</p>
        </div>

    )
} 
export default Cards;