import { BsHouseFill, BsPersonFill, BsQuestionCircleFill, BsPlayFill } from "react-icons/bs";
// import Card from "./Card"; 
import Ticket from './Ticket'
import { AiFillSetting } from "react-icons/ai";
import { useEffect, useState } from "react";

interface Email {
  email: string;
  title: string;
  summary: string;
  profilepic: string;
  status: string;
  tags: string[];
}

function Dashboard() {
    const [emails, setEmails] = useState<Email[]>([]);
    const loading = false;

    useEffect(() => {
        if(emails) {
            fetch('/api/emails', { 
                method: 'get', 
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                }
            }).then(res => res.json())
            .then(data => {
                console.log(data); 
                data.forEach((email: Email) => {
                    email.status = 'open';
                }); 
                setEmails(data);
            });
        }
    }, [loading]);
    

    return (
        <>
            <div>
                <div className="flex flex-col h-screen bg-gray-800 border-e-2 fixed myShadow top-0 text-white hidden md:flex">
                    <button className="m-2 p-1 border border-white rounded-2xl">
                        <BsHouseFill size='40px' color="white" className="m-auto"/>
                    </button>
                    <button className="m-2">
                        <BsPersonFill size='40px' color="white" className="m-auto"/>
                    </button>
                    <hr />
                    <button className="m-2 mt-auto">
                        <AiFillSetting size='40px' color="white" className="m-auto"/>
                    </button>
                    <button className="m-2">
                        <BsQuestionCircleFill size='40px' color="white" className="m-auto"/>
                    </button>
                </div>
                <div className="flex flex-col grow h-screen md:ml-16 bg-blue-50">
                    <h1 className="p-10 mr-auto text-2xl font-bold inline">Your workspace</h1>
                    
                    <ul className="flex ml-10 [&>li]:px-5 [&>li]:py-2 [&>li]:rounded-t-[80px_180px] [&>li]:bg-blue-200 [&>li]:shadow-lg">
                        <li className="active">
                            <a>Open</a>
                        </li>
                        <li>
                            <a>Waiting for response</a>
                        </li>
                        <li>
                            <a>Closed</a>
                        </li>
                        <li className="ml-auto mr-20">
                            <button className="ml-auto">Play <BsPlayFill className="inline"/></button>
                        </li>
                    </ul>
                    <ul className="mx-10">
                    {
                        emails.length === 0 ? 
                                            <></>:
                                            (emails.map((data, i) => {
                                                return <Ticket {...data} index={i} key={i}/>
                                            }))
                    }
                    </ul>
                </div>
            </div>
        </>
    )
}

export default Dashboard;