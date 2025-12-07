
type props = {
    title: string,
    summary: string,
    profilepic: string;
    score: number,
    tags: string[],
    index: number
}

function Alternate({direction, profilepic}: {direction: boolean, profilepic: string}) {
    const defaultPic = "https://lh3.googleusercontent.com/a/default-user=s100";
    if(direction)
        return (
        <>
            <div className="w-100"></div>
            {
                profilepic === "" ? <img src={defaultPic}  className='w-15 h-15 translate-x-8 translate-y-5 z-1 relative rounded-full'/> :
                                    <img src={profilepic}  className='w-15 h-15 translate-x-8 translate-y-5 z-1 relative rounded-full'/>
            }
            <div className={'w-1 bg-gray-300'} />
            <div className={'h-0.5 w-15 translate-y-13 bg-gray-300 inline'} />
        </>
        )
    else
        return (
        <>
            <div className={'h-0.5 w-15 translate-y-13 bg-gray-300 inline'} />
            <div className={'w-1 bg-gray-300 inline'} />
            {
                profilepic === "" ? <img src={defaultPic} className='w-15 h-15 translate-y-5 -translate-x-8 z-1 relative rounded-full'/>:
                                    <img src={profilepic} className='w-15 h-15 translate-y-5 -translate-x-8 z-1 relative rounded-full'/>
            }
            <div className="w-100"></div>
        </>
        )
}

function Card({title, summary, score, profilepic, tags, index} : props) {
    const direction = index % 2 === 0;
    tags.unshift(score.toString());

    return (
        <div className="flex text-left mx-auto">
            
        {direction ?  <Alternate direction profilepic={profilepic}/> : ''}
        <div className={`flex flex-col my-2 w-100 rounded-md myShadow border-2 border-gray-300 shadow-gray-500 bg-white`}>
            <div>
                <h1 className="text-xl font-medium p-2 rounded-t-md bg-gray-200">{title}</h1>
            </div>
            <div className="m-2">
                <p>{summary}</p>
            </div>
            <hr />
            <div className="m-2">
            {tags?.map((tag, i) => {
                return <span className="p-1 m-1 rounded-lg inline bg-gray-900 myShadow text-white" key={i}>{tag}</span>;
            })}
            </div>
        </div>
        {direction ? '' : <Alternate direction={direction} profilepic={profilepic}/>}
        </div>
    )
}

export default Card;