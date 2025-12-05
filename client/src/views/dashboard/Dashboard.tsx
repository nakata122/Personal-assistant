import { BsHouseFill, BsPersonFill, BsQuestionCircleFill } from "react-icons/bs";
import Card from "./Card"; 
import { AiFillSetting } from "react-icons/ai";
import { useEffect, useState } from "react";

interface Email {
  email: string;
  title: string;
  summary: string;
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
            <div className="flex">
                <div className="flex flex-col h-screen bg-gray-800 border-e-2 fixed myShadow top-0 text-white">
                    <div className="p-2 h-16" />
                    <button className="m-2 p-2 border border-white rounded-2xl shadow-sm shadow-white">
                        <BsHouseFill size='40px' color="white" className="m-auto"/>
                        <h1>Home</h1>
                    </button>
                    <button className="m-4">
                        <BsPersonFill size='40px' color="white" className="m-auto"/>
                        <h1>Clients</h1>
                    </button>
                    <hr />
                    <button className="m-4">
                        <AiFillSetting size='40px' color="white" className="m-auto"/>
                        <h1>Settings</h1>
                    </button>
                    <button className="m-4">
                        <BsQuestionCircleFill size='40px' color="white" className="m-auto"/>
                        <h1>Help</h1>
                    </button>

                </div>
                <div className="flex flex-col grow bg-white">
                    <h1 className="py-2 px-4 my-2 mx-auto rounded-lg bg-gray-900 text-white w-fit font-medium text-lg">Today</h1>
                    {
                        emails.length === 0 ? 
                                            <></>:
                                            (emails.map((data, i) => {
                                                return <Card {...data} index={i} />
                                            }))
                    }
                </div>
            </div>
        </>
    )
}

export default Dashboard;