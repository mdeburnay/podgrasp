// Dependencies
import React, { useState } from "react";
import axios from "axios";

// Styles
import logo from "./logo.svg";
import "./App.css";

function App() {
  const [url, setUrl] = useState<string>("");
  const [data, setData] = useState<string>("");

  const handlePasswordChange = (e: any) => {
    setUrl(e.target.value);
  };

  const handleSubmit = async (event: any) => {
    console.log("submitting");

    event.preventDefault();

    const data: any = await axios
      .post("http://localhost:9090/", url)
      .then((res) => {
        console.log(res.data.replaceAll("\n", ""));
      });
    setData(data);
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <form>
          <input
            type="url"
            name="URL"
            value={url}
            onChange={handlePasswordChange}
          />
          <button type="button" onClick={handleSubmit}>
            Login
          </button>
        </form>
        {data && <div>{data}</div>}
      </header>
    </div>
  );
}

export default App;
