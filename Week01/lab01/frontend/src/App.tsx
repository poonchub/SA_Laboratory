import React from "react";
import './App.css'

function App() {
  const studentId = "B6525163";
  const studentName = "Poonchub Nanawan";
  const teamNumber = "G06";

  return (
    <div className="App">
      <header className="App-header">
        <div>
          <img src="https://media0.giphy.com/media/v1.Y2lkPTc5MGI3NjExeW11MmpseXU1N2xxOGN2OXRiZDNibGp6cHJxZHB6dnMyYjh3d2szNyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/c5wbvuaVVLWzC/giphy.gif" alt="" />
        </div>
        <h1>สวัสดี, {studentId} {studentName} {teamNumber}!</h1>
      </header>
    </div>
  );
}

export default App
