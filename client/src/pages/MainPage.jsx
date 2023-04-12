import React from "react";
import { ENDPOINT } from "../App";

const Title = "iubrgre";
const Body = "uirbnnrtnk rtntjnrbjnrktnr tnhh rtn krkntj";

const MainPage = ({ data, mutate }) => {
  async function addTodo() {
    const newItem = {
      title: Title,
      body: Body,
    };
    await fetch(`${ENDPOINT}/api`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newItem),
    }).then((r) => r.json());

    mutate([...data, newItem]);
  }

  return <button onClick={addTodo}>Add</button>;
};

export default MainPage;
