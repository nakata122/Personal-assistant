import { BsHouseFill, BsPersonFill, BsQuestionCircleFill } from "react-icons/bs";
import Card from "./Card"; 
import { AiFillSetting } from "react-icons/ai";

function Dashboard() {
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
        },{
            username: 'John',
            title: 'Something BIG',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus.',
            tags: ['well formatted', 'formal', 'job offer']
        },{
            username: 'John',
            title: 'Something BIG',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus.',
            tags: ['well formatted', 'formal', 'job offer']
        },{
            username: 'John',
            title: 'Something BIG',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus.',
            tags: ['well formatted', 'formal', 'job offer']
        },{
            username: 'John',
            title: 'Something BIG',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus.',
            tags: ['well formatted', 'formal', 'job offer']
        },
    ];

  return (
    <>
        <div className="flex">
            <div className="flex flex-col h-screen bg-gray-800 border-e-2 myShadow fixed z-20 top-0 text-white">
                <div className="p-2 h-16 bg-gray-900">
                </div>
                <button className="m-2 p-2 border border-white rounded-2xl myShadow">
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
                    emails.map((data, i) => {
                        return <Card {...data} index={i} />
                    })
                }
            </div>
        </div>
    </>
  )
}

export default Dashboard;