import React, { useState, useEffect } from 'react';
const API = process.env.REACT_APP_API || "http://localhost:8080/api";

function App() {
  const [items, setItems] = useState([]);

  const loadItems = () =>
    fetch(`${API}/list`)
      .then(async (r) => {
        if (!r.ok) throw new Error(`Server error: ${r.status}`);
        const data = await r.json().catch(() => []);
        return Array.isArray(data) ? data : [];
      })
      .then(setItems)
      .catch((err) => {
        console.error("❌ loadItems failed:", err);
        setItems([]);
      });

  const add = () => fetch(`${API}/add?name=Item_${Date.now()}`).then(loadItems);
  const del = (id) => fetch(`${API}/delete?id=${id}`).then(loadItems);
  const load = () => fetch(`${API}/load`).then((r) => r.text()).then(alert);

  useEffect(loadItems, []);

  return (
    <div style={{ padding: 20 }}>
      <h2>Demo App</h2>
      <button onClick={add}>➕ Add</button>
      <button onClick={load} style={{ marginLeft: 10 }}>🔥 Load</button>
      <ul>
        {items.length === 0 ? (
          <li style={{ color: "#888" }}>Нет данных</li>
        ) : (
          items.map((it) => (
            <li key={it.id}>
              {it.name} <button onClick={() => del(it.id)}>❌</button>
            </li>
          ))
        )}
      </ul>
    </div>
  );
}

export default App;
