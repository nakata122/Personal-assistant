import { BsHouseFill, BsPersonFill, BsQuestionCircleFill } from "react-icons/bs";
// import Card from "./Card"; 
import Ticket from './Ticket'
import { AiFillSetting } from "react-icons/ai";
import { useEffect, useState } from "react";

interface Email {
  email: string;
  title: string;
  summary: string;
  profilepic: string;
  score: number;
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
                setEmails(data);
            });
        }
    }, [loading]);
    

    return (
        <>
            <div>
                <div className="flex flex-col h-screen bg-gray-800 border-e-2 fixed myShadow top-0 text-white hidden md:flex">
                    <div className="p-2 h-16" />
                    <button className="m-1 border border-white rounded-2xl shadow-sm shadow-white">
                        <BsHouseFill size='30px' color="white" className="m-auto"/>
                        <h1>Home</h1>
                    </button>
                    <button className="m-1">
                        <BsPersonFill size='30px' color="white" className="m-auto"/>
                        <h1>Contacts</h1>
                    </button>
                    <hr />
                    <button className="m-1">
                        <AiFillSetting size='30px' color="white" className="m-auto"/>
                        <h1>Settings</h1>
                    </button>
                    <button className="m-1">
                        <BsQuestionCircleFill size='30px' color="white" className="m-auto"/>
                        <h1>Help</h1>
                    </button>
                </div>
                <div className="flex flex-col grow min-h-[90vh] md:ml-[70px] bg-white">
                    {
                        emails.length === 0 ? 
                                            <></>:
                                            (emails.map((data, i) => {
                                                return <Ticket {...data} index={i} key={i}/>
                                            }))
                    }
                </div>
            </div>
        </>
    )
}

export default Dashboard;