import axios from "axios";
import { useState } from "react";

export default function Manage() {
  const [form, setForm] = useState({
    body: "",
  });

  const post = async (body: string) => {
    await axios.post("http://localhost:8000", { body });
  };

  return (
    <main>
      <h2 className="pb-2">New Note</h2>
      <div>
        <input
          type="text"
          value={form.body}
          onChange={(e) => setForm({ body: e.target.value })}
          className="bg-gray-50 mb-4 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
        />
      </div>
      <button
        onClick={() => post(form.body)}
        type="submit"
        className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-xs w-full sm:w-auto px-2 py-2 text-center"
      >
        Submit
      </button>
    </main>
  );
}
