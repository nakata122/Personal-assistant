
type props = {
    title: string,
    summary: string,
    profilepic: string;
    score: number,
    tags: string[],
    index: number
}

function Ticket({title, summary, score, profilepic, tags} : props) {
    const defaultPic = "https://lh3.googleusercontent.com/a/default-user=s100";

    return (
        <div className="flex flex-row text-left my-0.5 rounded-md myShadow border-2 border-gray-300 shadow-gray-500">
            <input type="checkbox" className="mx-2 my-auto size-4" />
            <div className={`flex flex-col w-full bg-white`}>
                <div className="flex text-xl font-medium p-2">
                    {
                    profilepic === "" ? <img src={defaultPic}  className='w-8 h-8 mr-2 inline rounded-full'/> :
                                        <img src={profilepic}  className='w-8 h-8 mr-2 inline rounded-full'/>
                    }
                    <h1 className="mx-2">{title}</h1>
                    {score}
                </div>
                <div className="m-1">
                    <p>{summary}</p>
                </div>
                <div className="m-0.5">
                {tags?.map((tag, i) => {
                    return <span className="p-0.5 m-0.5 rounded-lg inline bg-gray-900 myShadow text-white" key={i}>{tag}</span>;
                })}
                </div>
            </div>
        </div>
    )
}

export default Ticket;