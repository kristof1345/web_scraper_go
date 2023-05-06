import React from "react";
import axios from "axios";

const MainPage = ({ data, setData }) => {
  function postUrls(e) {
    let ret;
    e.preventDefault();
    let elems = document.getElementsByClassName("url");
    let items = document.getElementsByClassName("item");
    let urls = [...elems].map((elem) => elem.value);
    let htmlELems = [...items].map((elem) => elem.value);
    const newItem = {
      url1: urls[0],
      item1: htmlELems[0],
      url2: urls[1],
      item2: htmlELems[1],
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
        <div>
          <input type="text" className="url" />
          <input type="text" className="item" />
        </div>
        <div>
          <input type="text" className="url" />
          <input type="text" className="item" />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default MainPage;
