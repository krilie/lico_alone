import React from "react";
import "./Photos.less"
import {connect} from "react-redux";
import actions from "../../redux/actions/ActionCreator";

class Photos extends React.Component {

    todoChange = (event) => {
        this.props.toggleComplete(event.target.value);
    };

    getTodos() {
        return this.props.todos.map((todo, index) => {
            return (<li key={index}>
                <input type="checkbox"
                       value={todo.id}
                       onClick={this.todoChange}
                       checked={todo.isComplete}/> {
                todo.isComplete
                    ? <del>{todo.title}</del>
                    : <span>{todo.title}</span>
            }
                <button type="button" data-id={todo.id}>删除</button>
            </li>);
        });
    }

    render() {
        return (
            <div>
                <ul>
                    {this.getTodos()}
                </ul>
            </div>);
    }
}

export default Photos = connect((state) => ({...state}),actions)(Photos);