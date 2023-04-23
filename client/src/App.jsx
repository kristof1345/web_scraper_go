import { useState, useEffect } from "react";
import MainPage from "./pages/MainPage";
import Post from "./comps/post";

function App() {
  const [data, setData] = useState(
    JSON.parse(localStorage.getItem("data")) || []
  );

  useEffect(() => {
    localStorage.setItem("data", JSON.stringify(data));
  }, [data]);

  console.log(data);

  return (
    <div className="App">
      <h1>Scraper</h1>
      <MainPage data={data} setData={setData} />
      <div className="container">
        {data.length > 0
          ? data.map((ret, i) => <Post post={ret} key={i} />)
          : null}
      </div>
    </div>
  );
}

export default App;
