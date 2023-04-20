import React from "react";
import axios from "axios";

const MainPage = ({ data, setData }) => {
  function postUrls(e) {
    let ret;
    e.preventDefault();
    let elems = document.getElementsByClassName("url");
    let urls = [...elems].map((elem) => elem.value);
    const newItem = {
      url1: urls[0],
      url2: urls[1],
    };
    axios.post("http://localhost:4000/api", newItem).then((r) => {
      ret = r.data;
      setData(ret);
    });
  }

  // function test(e) {
  //   e.preventDefault();
  //   let elems = document.getElementsByClassName("url");
  //   let urls = [...elems].map((elem) => elem.value);
  //   console.log(urls);
  // }

  return (
    <div>
      <form onSubmit={(e) => postUrls(e)}>
        <input type="text" className="url" />
        <input type="text" className="url" />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default MainPage;
