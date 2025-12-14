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
            <button className="btn" onClick={() => setView("create")}>Create Room</button>
            <button className="btn-secondary" onClick={() => setView("join")}>Join Room</button>
            <button className="btn-outline" onClick={() => setView("about")}>About</button>
          </>
        )}


        {view === "create" && (
          <>
            <h2 className="text-xl font-semibold">Create Room</h2>
            <input className="input" placeholder="Room name" value={name} onChange={e => setName(e.target.value)} />
            <textarea className="input" placeholder="Description (optional)" value={description} onChange={e => setDescription(e.target.value)} />
            <button className="btn">Create & Get Link</button>
            <button className="btn-ghost" onClick={() => setView("landing")}>Back</button>
          </>
        )}


        {view === "join" && (
          <>
            <h2 className="text-xl font-semibold">Join Room</h2>
            <input className="input" placeholder="Room ID" value={roomId} onChange={e => setRoomId(e.target.value)} />
            <button className="btn">Join</button>
            <button className="btn-ghost" onClick={() => setView("landing")}>Back</button>
          </>
        )}


        {view === "about" && (
          <>
            <h2 className="text-xl font-semibold">About</h2>
            <p className="text-sm text-gray-600">
              Organize Secret Santa gift exchanges directly inside Telegram.
            </p>
            <button className="btn-ghost" onClick={() => setView("landing")}>Back</button>
          </>
        )}
      </div>
    </main>
  );
}