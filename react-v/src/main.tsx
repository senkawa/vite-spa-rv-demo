import ReactDOM from "react-dom/client";
import App from "./App";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Manage from "./pages/Manage";
import Notes from "./pages/Notes";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<App />}>
        <Route path="manage" element={<Manage />} />
        <Route path="notes" element={<Notes />} />
      </Route>
    </Routes>
  </BrowserRouter>
);
