"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";

export default function Login() {
  const router = useRouter();
  const [user, setUser] = useState({ username: "", password: "" });

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://localhost:8000/login", user);
      localStorage.setItem("token", res.data.token);
      router.push("/dashboard");
    } catch {
      alert("Invalid credentials");
    }
  };

  return (
    <div className="flex justify-center items-center h-screen">
      <form className="bg-white p-6 rounded shadow-lg" onSubmit={handleLogin}>
        <h2 className="text-xl font-bold mb-4">Login</h2>
        <input className="border p-2 w-full mb-2" placeholder="Username" value={user.username} onChange={(e) => setUser({ ...user, username: e.target.value })} required />
        <input className="border p-2 w-full mb-2" type="password" placeholder="Password" value={user.password} onChange={(e) => setUser({ ...user, password: e.target.value })} required />
        <button className="bg-blue-500 text-white p-2 w-full rounded" type="submit">Login</button>
      </form>
    </div>
  );
}
