"use client";


import { useEffect, useState } from "react";


export default function Home() {
  const [view, setView] = useState<"landing" | "create" | "join" | "about">("landing");
  const [roomId, setRoomId] = useState("");
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");


  useEffect(() => {
    if (typeof window !== "undefined" && (window as any).Telegram?.WebApp) {
      const tg = (window as any).Telegram.WebApp;
      tg.ready();
      tg.expand();
    }
  }, []);


  return (
    <main className="min-h-screen flex items-center justify-center p-4 bg-gradient-to-br from-red-50 to-green-50">
      <div className="w-full max-w-md bg-white rounded-2xl shadow-xl p-6 space-y-4">
        {view === "landing" && (
          <>
            <h1 className="text-3xl font-bold text-center">🎅 Secret Santa</h1>
            <button className="w-full bg-red-600 text-white py-2 rounded-xl font-medium;" onClick={() => setView("create")}>Create Room</button>
            <button className="w-full bg-green-600 text-white py-2 rounded-xl font-medium;" onClick={() => setView("join")}>Join Room</button>
            <button className="w-full border border-gray-300 py-2 rounded-xl" onClick={() => setView("about")}>About</button>
          </>
        )}


        {view === "create" && (
          <>
            <h2 className="text-xl font-semibold">Create Room</h2>
            <input className="border rounded-xl px-3 py-2" placeholder="Room name" value={name} onChange={e => setName(e.target.value)} />
            <textarea className="border rounded-xl px-3 py-2" placeholder="Description (optional)" value={description} onChange={e => setDescription(e.target.value)} />
            <button className="w-full bg-red-600 text-white py-2 rounded-xl font-medium;">Create & Get Link</button>
            <button className="w-full text-gray-500 py-2" onClick={() => setView("landing")}>Back</button>
          </>
        )}


        {view === "join" && (
          <>
            <h2 className="text-xl font-semibold">Join Room</h2>
            <input className="border rounded-xl px-3 py-2" placeholder="Room ID" value={roomId} onChange={e => setRoomId(e.target.value)} />
            <button className="w-full bg-red-600 text-white py-2 rounded-xl font-medium;">Join</button>
            <button className="w-full text-gray-500 py-2" onClick={() => setView("landing")}>Back</button>
          </>
        )}


        {view === "about" && (
          <>
            <h2 className="text-xl font-semibold">About</h2>
            <p className="text-sm text-gray-600">
              Organize Secret Santa gift exchanges directly inside Telegram.
            </p>
            <button className="w-full text-gray-500 py-2" onClick={() => setView("landing")}>Back</button>
          </>
        )}
      </div>
    </main>
  );
}