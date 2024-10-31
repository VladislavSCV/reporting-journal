async function registerUser(login, password) {
  const response = await fetch('http://localhost:8080/api/auth/registration', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ login, password })
  });
  return response.json();
}

async function loginUser(login, password) {
  const response = await fetch('http://localhost:8080/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ login, password })
  });
  if (response.ok) {
    const data = await response.json();
    localStorage.setItem('token', data.token); // Сохраняем JWT токен
    return data;
  } else {
    throw new Error('Login failed');
  }
}

async function getCurrentUser() {
  const token = localStorage.getItem('token');
  const response = await fetch('http://localhost:8080/api/auth', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${token}` }
  });
  return response.json();
}

async function getUsers() {
  const response = await fetch('http://localhost:8080/api/user', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function getUserById(userId) {
  const response = await fetch(`http://localhost:8080/api/user/${userId}`, {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteUser(userId) {
  const response = await fetch(`http://localhost:8080/api/user/${userId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function createRole(value) {
  const response = await fetch('http://localhost:8080/api/role', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ value })
  });
  return response.json();
}

async function getRoles() {
  const response = await fetch('http://localhost:8080/api/role', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteRole(roleId) {
  const response = await fetch(`http://localhost:8080/api/role/${roleId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function addGroup(name, body) {
  const response = await fetch('http://localhost:8080/api/groups', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ name, body })
  });
  return response.json();
}

async function getGroups() {
  const response = await fetch('http://localhost:8080/api/groups', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteGroup(groupId) {
  const response = await fetch(`http://localhost:8080/api/groups/${groupId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateGroup(groupId, name, body) {
  const response = await fetch(`http://localhost:8080/api/groups/${groupId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ name, body })
  });
  return response.json();
}

async function addNote(title, body, groupId) {
  const response = await fetch('http://localhost:8080/api/notes', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ title, body, groupId })
  });
  return response.json();
}

async function getNotes() {
  const response = await fetch('http://localhost:8080/api/notes', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteNote(noteId) {
  const response = await fetch(`http://localhost:8080/api/notes/${noteId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateNote(noteId, title, body) {
  const response = await fetch(`http://localhost:8080/api/notes/${noteId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ title, body })
  });
  return response.json();
}

async function addSchedule(groupId, dayOfWeek, subject, teacher) {
  const response = await fetch('http://localhost:8080/api/schedule', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ groupId, dayOfWeek, subject, teacher })
  });
  return response.json();
}

async function getSchedules() {
  const response = await fetch('http://localhost:8080/api/schedule', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteSchedule(scheduleId) {
  const response = await fetch(`http://localhost:8080/api/schedule/${scheduleId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateSchedule(scheduleId, groupId, dayOfWeek, subject, teacher) {
  const response = await fetch(`http://localhost:8080/api/schedule/${scheduleId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ groupId, dayOfWeek, subject, teacher })
  });
  return response.json();
}

async function addStudent(name, groupId, role) {
  const response = await fetch('http://localhost:8080/api/students', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ name, groupId, role })
  });
  return response.json();
}

async function getStudents() {
  const response = await fetch('http://localhost:8080/api/students', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteStudent(studentId) {
  const response = await fetch(`http://localhost:8080/api/students/${studentId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateStudent(studentId, name, groupId, role) {
  const response = await fetch(`http://localhost:8080/api/students/${studentId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ name, groupId, role })
  });
  return response.json();
}
