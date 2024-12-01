import {useNavigate} from "react-router-dom";

export async function registerUser(first_name, middle_name, last_name, login, password, role_id, group_id) {
  console.log(first_name, middle_name, last_name, login, password, role_id, group_id);
  role_id = Number(role_id)
  group_id = Number(group_id)

  try {
    const response = await fetch('/api/auth/registration', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ first_name, middle_name, last_name, login, password, role_id, group_id }),
    });

    // Проверка статуса ответа
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Ошибка регистрации');
    }

    // Извлечение данных
    const res = await response.json();

    return res; // Возврат данных
  } catch (error) {
    console.error('Ошибка регистрации:', error.message);
    throw error; // Проброс ошибки для обработки в вызывающем коде
  }
}

export async function loginUser(login, password, navigate) {
  try {
    const response = await fetch('/api/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ login, password }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Login failed');
    }

    const data = await response.json();

    if (data.token) {
      localStorage.setItem("token", data.token);
      localStorage.setItem("user_id", data.user.id);
      if (data.user.role_id) {
        localStorage.setItem("group_id", null);
      }
      localStorage.setItem("group_id", data.user.group_id);

      console.log("Login successful:", data);
      navigate("/mainPage");
    } else {
      throw new Error('Invalid response data: missing token, id, or group_id');
    }
  } catch (error) {
    console.error('Error during login:', error);
    alert(error.message || 'Login failed. Please try again.');
    throw error; // Если требуется переброс ошибки выше
  }

}


export function Decode(Response) {
  const database64 = Response.json();
  const data = JSON.parse(atob(database64));
  return data
}

async function VerifyTokenAndGetId(token) {
    const response = await fetch('/api/auth/verify', {
      method: 'POST',
    })

    if (!response.ok) {
      const errorData = await response.json();
      console.error('Error verifying token:', errorData);
      return false;
    }

    const data = await response.json();
    return data.id
}


async function getCurrentUser() {
  const token = localStorage.getItem('token');
  const response = await fetch('http://localhost:8000/api/auth', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${token}` }
  });
  return response.json();
}

async function getUsers() {
  const response = await fetch('http://localhost:8000/api/user', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function getUserById(userId) {
  const response = await fetch(`http://localhost:8000/api/user/${userId}`, {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

export async function deleteUser(userId) {
  try {
    const response = await fetch(`/api/user/${userId}`, {
      method: 'DELETE',
      headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` },
    });
    if (!response.ok) {
      const errorData = await response.json();
      console.error('Error deleting user:', errorData);
      return false;
    }
    return true;
  } catch (error) {
    console.error('Request failed:', error);
    return false;
  }
}

export async function createRole(value) {
  const response = await fetch('/api/role', {
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
  const response = await fetch('http://localhost:8000/api/role', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteRole(roleId) {
  const response = await fetch(`http://localhost:8000/api/role/${roleId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

export async function addGroup(name) {
  const response = await fetch('/api/group', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ name })
  });
  return response.json();
}

export async function getGroups() {
  const response = await fetch('/api/group', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteGroup(groupId) {
  const response = await fetch(`http://localhost:8000/api/groups/${groupId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateGroup(groupId, name, body) {
  const response = await fetch(`http://localhost:8000/api/groups/${groupId}`, {
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
  const response = await fetch('http://localhost:8000/api/notes', {
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
  const response = await fetch('http://localhost:8000/api/notes', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteNote(noteId) {
  const response = await fetch(`http://localhost:8000/api/notes/${noteId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateNote(noteId, title, body) {
  const response = await fetch(`http://localhost:8000/api/notes/${noteId}`, {
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
  const response = await fetch('http://localhost:8000/api/schedule', {
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
  const response = await fetch('http://localhost:8000/api/schedule', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteSchedule(scheduleId) {
  const response = await fetch(`http://localhost:8000/api/schedule/${scheduleId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateSchedule(scheduleId, groupId, dayOfWeek, subject, teacher) {
  const response = await fetch(`http://localhost:8000/api/schedule/${scheduleId}`, {
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
  const response = await fetch('http://localhost:8000/api/students', {
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
  const response = await fetch('http://localhost:8000/api/students', {
    method: 'GET',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.json();
}

async function deleteStudent(studentId) {
  const response = await fetch(`http://localhost:8000/api/students/${studentId}`, {
    method: 'DELETE',
    headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
  });
  return response.ok;
}

async function updateStudent(studentId, name, groupId, role) {
  const response = await fetch(`http://localhost:8000/api/students/${studentId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ name, groupId, role })
  });
  return response.json();
}
