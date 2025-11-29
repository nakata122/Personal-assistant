import { useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { BsXCircleFill, BsPersonFill } from "react-icons/bs";

type ModalProps = {
    toggle: () => void
};

function Modal({ toggle } : ModalProps) {
    const [username, setUsername] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    function handleLogin() {

    }

    return (<div className="flex absolute top-0 left-0 h-screen w-screen backdrop-blur-sm z-20">
            <div className="m-auto flex flex-col bg-white rounded-2xl text-black items-center w-md h-1/2">
                <BsXCircleFill onClick={toggle} className="w-6 h-6 m-2 relative ml-auto cursor-pointer" />

                <h1 className="text-4xl font-bold mt-10">Login</h1>
                <form onSubmit={handleLogin}  className="flex flex-col w-full" >
                    <button type="submit" className="flex items-center p-2 border-2 rounded-2xl border-gray-900 m-5 hover:bg-gray-200 cursor-pointer"> 
                        <FcGoogle className="w-6 h-6 inline mr-auto"/> 
                        <span className="m-2 font-bold mr-auto">Continue with google</span>
                    </button>
                    
                    <button type="submit" className="flex items-center p-2 border-2 rounded-2xl border-gray-900 m-5 hover:bg-gray-200 cursor-pointer"> 
                        <BsPersonFill className="w-6 h-6 inline mr-auto"/> 
                        <span className="m-2 font-bold mr-auto">Continue as guest</span>
                    </button>
                </form>
            </div>
        </div>);
}

export default Modal;