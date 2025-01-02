document.getElementById('addTaskBtn').addEventListener('click', function () {
  const taskInput = document.getElementById('taskInput');
  const taskList = document.getElementById('taskList');
  const noTasksMsg = document.getElementById('noTasksMsg');
  const statusMsg = document.getElementById('statusMsg');
  const pendingCount = document.getElementById('pendingCount');

  const task = taskInput.value.trim();

  if (task) {
    const li = document.createElement('li');
    li.setAttribute('draggable', 'true'); // Make the task draggable

    const taskText = document.createElement('span');
    taskText.textContent = task;

    const completeBtn = document.createElement('button');
    completeBtn.textContent = 'Complete';
    completeBtn.style.marginLeft = '10px';
    completeBtn.style.backgroundColor = '#17a2b8';
    completeBtn.style.color = 'white';
    completeBtn.style.border = 'none';
    completeBtn.style.borderRadius = '4px';
    completeBtn.style.padding = '4px 8px';
    completeBtn.style.cursor = 'pointer';

    completeBtn.addEventListener('click', function () {
      if (!taskText.style.textDecoration || taskText.style.textDecoration === 'none') {
        taskText.style.textDecoration = 'line-through';
        taskText.style.color = '#6c757d'; // Optional: Change text color to indicate completion
        statusMsg.textContent = 'Task marked as completed!';
        statusMsg.style.color = '#17a2b8';
        updatePendingCount();
      } else {
        taskText.style.textDecoration = 'none';
        taskText.style.color = ''; // Reset text color
        statusMsg.textContent = 'Task unmarked as completed!';
        statusMsg.style.color = '#ffc107';
        updatePendingCount();
      }

      setTimeout(() => {
        statusMsg.textContent = '';
      }, 3000);
    });

    const editBtn = document.createElement('button');
    editBtn.textContent = 'Edit';
    editBtn.style.marginLeft = '10px';
    editBtn.style.backgroundColor = '#007bff';
    editBtn.style.color = 'white';
    editBtn.style.border = 'none';
    editBtn.style.borderRadius = '4px';
    editBtn.style.padding = '4px 8px';
    editBtn.style.cursor = 'pointer';

    editBtn.addEventListener('click', function () {
      const newTask = prompt('Edit your task:', taskText.textContent);
      if (newTask !== null) {
        taskText.textContent = newTask.trim();
        statusMsg.textContent = 'Task was updated successfully!';
        statusMsg.style.color = '#ffc107';

        setTimeout(() => {
          statusMsg.textContent = '';
        }, 3000);
      }
    });

    const deleteBtn = document.createElement('button');
    deleteBtn.textContent = 'Delete';
    deleteBtn.style.marginLeft = '10px';
    deleteBtn.style.backgroundColor = '#dc3545';
    deleteBtn.style.color = 'white';
    deleteBtn.style.border = 'none';
    deleteBtn.style.borderRadius = '4px';
    deleteBtn.style.padding = '4px 8px';
    deleteBtn.style.cursor = 'pointer';

    deleteBtn.addEventListener('click', function () {
      li.remove();
      statusMsg.textContent = 'Task was deleted successfully!';
      statusMsg.style.color = '#dc3545';
      updatePendingCount();
      if (taskList.children.length === 0) {
        noTasksMsg.style.display = 'block';
      }

      setTimeout(() => {
        statusMsg.textContent = '';
      }, 3000);
    });

    // Drag-and-Drop Event Handlers
    li.addEventListener('dragstart', function (e) {
      e.dataTransfer.setData('text/plain', li.dataset.index);
      li.classList.add('dragging');
    });

    li.addEventListener('dragover', function (e) {
      e.preventDefault();
      const dragging = document.querySelector('.dragging');
      if (dragging) {
        taskList.insertBefore(dragging, li.nextSibling);
      }
    });

    li.addEventListener('drop', function (e) {
      e.preventDefault();
      li.classList.remove('dragging');
    });

    li.addEventListener('dragend', function () {
      li.classList.remove('dragging');
    });

    li.appendChild(taskText);
    li.appendChild(completeBtn);
    li.appendChild(editBtn)
    li.appendChild(deleteBtn);
    taskList.appendChild(li);

    taskInput.value = '';
    noTasksMsg.style.display = 'none';
    statusMsg.textContent = 'Task was added successfully!';
    statusMsg.style.color = '#28a745';

    updatePendingCount();

    setTimeout(() => {
      statusMsg.textContent = '';
    }, 3000);
  }
});

document.getElementById('taskInput').addEventListener('input', function () {
  const taskList = document.getElementById('taskList');
  const noTasksMsg = document.getElementById('noTasksMsg');

  if (taskList.children.length === 0) {
    noTasksMsg.style.display = 'block';
  }
});

function updatePendingCount() {
  const taskList = document.getElementById('taskList');
  const pendingCount = document.getElementById('pendingCount');
  let count = 0;

  Array.from(taskList.children).forEach((task) => {
    const taskText = task.querySelector('span');
    if (!taskText.style.textDecoration || taskText.style.textDecoration === 'none') {
      count++;
    }
  });

  pendingCount.textContent = `Pending Tasks: ${count}`;
}




  