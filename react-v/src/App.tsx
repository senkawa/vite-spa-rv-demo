import "./App.css";
import { Link, Outlet } from "react-router-dom";

function App() {
  return (
    <div className="m-auto w-60 text-sm pt-5">
      <nav className="border-b-2 pb-3">
        <Link to="/manage">Manage</Link> | <Link to="/notes">Notes</Link>
      </nav>
      <Outlet />
    </div>
  );
}

export default App;
