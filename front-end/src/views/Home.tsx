import { useState } from "react";
import Login from "../shared/Login";

function Home() {
  const [login, setLogin] = useState<boolean>(false);

  function toggleLogin() {
    setLogin(!login);
  };

  return (
    <div className="flex flex-col text-white h-screen w-screen bg-linear-to-b from-sky-600 to-blue-300">
        <h1 className="title mx-20 my-5 font-bold">
            Upload and Manage your clients with an email.
        </h1>
        <p className="text-2xl w-2/3 ml-auto mr-auto my-5">
            This app loads clients automatically from the email with one click. Creates unique client score, monitor behavior, detects red flags. Additionaly you can visualize, summarize emails on an interactive timeline with AI.
        </p>
        <div className="flex flex-row justify-center font-medium">
            <button className="p-3 m-3 rounded-2xl bg-white text-4xl text-blue-950 shadow-2xl hover:bg-gray-100">
                Try Demo
            </button>
            <button onClick={toggleLogin} className="p-3 m-3 rounded-2xl bg-white text-4xl text-blue-950 shadow-2xl hover:bg-gray-100">
                Start Now
            </button>
        </div>
        {login ? <Login toggle={toggleLogin}/> : null}
    </div>
  )
}

export default Home;