import { BsCircleFill } from "react-icons/bs";

type props = {
    title: string,
    summary: string,
    tags: string[],
    score: number,
    index: number
}

function Alternate({direction}: {direction: boolean}) {
    if(direction)
        return (
        <>
            <div className="w-100"></div>
            <BsCircleFill size='60px' className='translate-x-8  translate-y-5 z-1 relative'/>
            <div className={'w-1 bg-gray-300'} />
            <div className={'h-0.5 w-15 translate-y-13 bg-gray-300 inline'} />
        </>
        )
    else
        return (
        <>
            <div className={'h-0.5 w-15 translate-y-13 bg-gray-300 inline'} />
            <div className={'w-1 bg-gray-300 inline'} />
            <BsCircleFill size='60px' className='translate-y-5 -translate-x-8 z-1 relative'/>
            <div className="w-100"></div>
        </>
        )
}

function Card({title, summary, score, tags, index} : props) {
    const direction = index % 2 === 0;
    tags.unshift(score.toString());

  return (
    <div className="flex text-left mx-auto">
        
    {direction ?  <Alternate direction /> : ''}
    <div className={`flex flex-col my-2 w-100 rounded-md myShadow border-2 border-gray-300 shadow-gray-500 bg-white`}>
        <div>
            <h1 className="text-xl font-medium p-2 rounded-t-md bg-gray-200">{title}</h1>
        </div>
        <div className="m-2">
            <p>{summary}</p>
        </div>
        <hr />
        <div className="m-2">
        {tags?.map(tag => {
            return <span className="p-1 m-1 rounded-lg inline bg-gray-900 myShadow text-white">{tag}</span>;
        })}
        </div>
    </div>
    {direction ? '' : <Alternate direction={direction} />}
    </div>
  )
}

export default Card;