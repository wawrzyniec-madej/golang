const SERVER_ADDRESS = 'http://0.0.0.0:8080';

const fetchTodos = async () => {
    const response = await fetch(SERVER_ADDRESS+'/api/todos');
    const todos = await response.json();
    return todos;
}

const addTodo = async (message) => {
    const response = await fetch(
        SERVER_ADDRESS+'/api/todos',
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                message
            })
        }
    );

    return response;
}

const clearTodos = async () => {
    const response = await fetch(
        SERVER_ADDRESS+'/api/todos',
        {
            method: 'DELETE'
        }
    );

    return response;
}

const reloadTodos = () => {

    const todoContainerElement = document.getElementById("todo-container");
    const todoTemplate = document.getElementById("todo-template");

    todoContainerElement.querySelectorAll(".todo-message").forEach(element => {
        element.remove()
    })

    fetchTodos().then(todos => {

        todos.forEach(todo => {

            const clone = todoTemplate.content.cloneNode(true);

            clone.querySelectorAll(".todo-message")[0].innerHTML = todo

            todoContainerElement.appendChild(clone)
        })
    })
}

window.onload = () => {

    const todoAddInput = document.getElementById("todo-add-input");
    const todoAddForm = document.getElementById("todo-add-form");
    const todoButtonClearElement = document.getElementById("todo-button-clear");

    reloadTodos()

    todoAddForm.addEventListener("submit", (event) => {

        event.preventDefault();
        addTodo(todoAddInput.value).then(response => {

            if (response.status != 200) {
                console.error("Submit did not create new todo")
                return
            }

            reloadTodos()
            todoAddInput.value = null
        })
    });

    todoButtonClearElement.addEventListener("click", event => {
        clearTodos().then(response => {
            if (response.status != 200) {
                console.error("Clearing todos did not work")
                return
            }

            reloadTodos()
        })
    });
}
