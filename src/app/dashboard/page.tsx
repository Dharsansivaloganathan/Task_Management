"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";
import { io } from "socket.io-client";

const API_BASE_URL = "http://localhost:8000";
const socket = io(API_BASE_URL);

export default function Dashboard() {
  const router = useRouter();
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState({ title: "", description: "" });

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) return router.push("/login");
    fetchTasks();

    socket.on("task_update", fetchTasks);

    return () => {
      socket.off("task_update", fetchTasks);
    };
  }, []);

  const fetchTasks = async () => {
    try {
      const res = await axios.get(`${API_BASE_URL}/tasks`);
      setTasks(res.data);
    } catch (error) {
      console.error("Error fetching tasks:", error);
    }
  };

  const handleCreateTask = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await axios.post(`${API_BASE_URL}/tasks`, newTask, {
        headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
      });
      setNewTask({ title: "", description: "" });
      fetchTasks();
    } catch (error) {
      console.error("Error creating task:", error);
    }
  };

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Task Dashboard</h1>
      <form onSubmit={handleCreateTask} className="mb-4">
        <input
          className="border p-2 mr-2"
          placeholder="Title"
          value={newTask.title}
          onChange={(e) => setNewTask({ ...newTask, title: e.target.value })}
          required
        />
        <input
          className="border p-2 mr-2"
          placeholder="Description"
          value={newTask.description}
          onChange={(e) => setNewTask({ ...newTask, description: e.target.value })}
          required
        />
        <button className="bg-blue-500 text-white p-2 rounded" type="submit">
          Add Task
        </button>
      </form>
      <ul>
        {tasks.map((task: any) => (
          <li key={task.id} className="border p-3 mb-2 rounded">
            <strong>{task.title}</strong>: {task.description} ({task.status})
          </li>
        ))}
      </ul>
    </div>
  );
}