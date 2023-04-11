import { useState } from "react";
import useSWR from "swr";
import MainPage from "./pages/MainPage";

// export const ENDPOINT = "http://127.0.0.1:5173";
export const ENDPOINT = "http://localhost:4000";

const fetcher = (url) => fetch(`${ENDPOINT}${url}`).then((r) => r.json());

function App() {
  const { data, mutate } = useSWR("/api/todos", fetcher);

  console.log(JSON.stringify(data));

  return (
    <div className="App">
      <div>{JSON.stringify(data)}</div>
      <h1>Vite + React</h1>
      <MainPage data={data} mutate={mutate} />
    </div>
  );
}

export default App;
