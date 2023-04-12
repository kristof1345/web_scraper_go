import { useState } from "react";
import useSWR from "swr";
import MainPage from "./pages/MainPage";
import Post from "./comps/post";

// export const ENDPOINT = "http://127.0.0.1:5173";
export const ENDPOINT = "http://localhost:4000";

const fetcher = (url) => fetch(`${ENDPOINT}${url}`).then((r) => r.json());

function App() {
  const { data, mutate } = useSWR("/get", fetcher);

  // const retData = JSON.stringify(data);
  let obj = {};

  if (data) {
    obj = Object.entries(data);
    console.log(obj);
  }

  return (
    <div className="App">
      <h1>Scraper</h1>
      <MainPage data={data} mutate={mutate} />
      <div className="container">
        {obj.length > 0
          ? obj.map((ret, i) => <Post post={ret[1]} key={i} />)
          : null}
      </div>
    </div>
  );
}

export default App;
