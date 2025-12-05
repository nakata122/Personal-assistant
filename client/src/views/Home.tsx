import { useState } from "react";
import Login from "../shared/Login";
import RevealAnimation from "../shared/RevealAnimation";
import { BsArrowDownShort, BsCloudArrowDownFill, BsDatabaseFillAdd, BsEnvelopeFill } from "react-icons/bs";
import Card from "./dashboard/Card";

function Home() {
  const [login, setLogin] = useState<boolean>(false);
  const emails = [
        {
            username: 'John',
            title: 'Something BIG',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus.',
            tags: ['well formatted', 'formal', 'job offer']
        },{
            username: 'John',
            title: 'Something BIG',
            summary: 'Lorem ipt dignissimos cupiditate doloremque? Asperiores possimus.',
            tags: ['well formatted', 'formal', 'job offer']
        },{
            username: 'John',
            title: 'Something BIG',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus.',
            tags: ['well formatted', 'formal', 'job offer']
        }
    ];

  function toggleLogin() {
    setLogin(!login);
  };

  function handleGuest() {
    fetch('/api/ping', { 
        method: 'get', 
        mode: 'no-cors',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json'
        }
    }).then(res => {
      console.log(res);
      return res.json();
    }).then(data => console.log(data));
    
  }

  return (
  <>
    <div className="flex flex-col text-white h-screen bg-linear-to-b from-sky-600 to-blue-200">
        {login ? <Login toggle={toggleLogin}/> : null}
        <h1 className="title mx-20 my-5 font-bold">
            Zero-setup customer identity recognition
        </h1>
        <p className="text-2xl w-1/3 ml-auto mr-auto my-5">
            An NLP and AI system that analyzes every email across your company and builds a real-time customer profile.
        </p>
        <div className="flex flex-row justify-center font-medium">
            <button onClick={handleGuest} className="p-3 m-3 rounded-2xl bg-white text-4xl text-blue-950 shadow-2xl hover:bg-gray-100">
                Try Demo
            </button>
            <button onClick={toggleLogin} className="p-3 m-3 rounded-2xl bg-white text-4xl text-blue-950 shadow-2xl hover:bg-gray-100">
                Start Now
            </button>
        </div>
        <BsArrowDownShort size='50px' className="animate-bounce m-auto"/>
    </div>
      <h1 className="title mx-20 my-20 font-bold text-gray-900">
          How it works?
      </h1>
    <div className="flex flex-row flex-1 justify-center text-white bg-linear-to-b from-blue-200 to-sky-600">
      <RevealAnimation>
      <div className="h-80 bg-gray-900 rounded-2xl w-70 mx-5">
        <h1 className="m-5 font-bold text-3xl">Login with email</h1>
        <BsEnvelopeFill size='120px' color="black" className="mx-auto mt-10 p-3 bg-white rounded-full shadow-2xl"/>
        <p className="m-5 text-xl">Using one or multiple emails</p>
      </div>
      </RevealAnimation>

      <RevealAnimation>
      <div className="h-80 bg-gray-900 rounded-2xl w-70 mx-5">
        <h1 className="m-5 font-bold text-3xl">Extract data</h1>
        <BsDatabaseFillAdd size='120px' color="black" className="mx-auto my-5 p-3 bg-white rounded-full shadow-2xl"/>
        <p className="m-5 text-xl">Parse all the emails, extract, store, organize client information</p>
      </div>
      </RevealAnimation>

      <RevealAnimation>
      <div className="h-80 bg-gray-900 rounded-2xl w-70 mx-5">
        <h1 className="m-5 font-bold text-3xl ">Download</h1>
        <BsCloudArrowDownFill size='120px' color="black" className="mx-auto p-3 bg-white rounded-full shadow-2xl"/>
        <p className="m-5 text-xl">Access data via JSON</p>
      </div>
      </RevealAnimation>
      <div className="h-100"></div>
    </div>
    
    <div className="bg-linear-to-b from-sky-600 to-blue-200 pb-20">
      <h1 className="title mx-20 mb-20 font-bold text-white">
          Timeline summary
      </h1>

      <h1 className="py-2 px-4 my-2 mx-auto rounded-lg bg-gray-900 text-white w-fit font-medium text-lg">Today</h1>
      {
          emails.map((data, i) => {
              return <RevealAnimation><Card {...data} index={i} /></RevealAnimation>
          })
      }
    </div>
  </>
  )
}

export default Home;