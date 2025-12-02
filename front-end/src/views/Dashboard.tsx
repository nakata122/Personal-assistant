import { BsHouseFill, BsPersonFill } from "react-icons/bs";

function Dashboard() {
    const emails = [
        {
            username: 'John',
            title: 'title',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus, ratione hic quo obcaecati magni modi eum, ipsa saepe velit architecto.'
        },
        {
            username: 'John',
            title: 'title',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus, ratione hic quo obcaecati magni modi eum, ipsa saepe velit architecto.'
        },
        {
            username: 'John',
            title: 'title',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus, ratione hic quo obcaecati magni modi eum, ipsa saepe velit architecto.'
        },
        {
            username: 'John',
            title: 'title',
            summary: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem alias harum deserunt at est dignissimos cupiditate doloremque? Asperiores possimus, ratione hic quo obcaecati magni modi eum, ipsa saepe velit architecto.'
        }
    ];

  return (
    <>
        <div className="flex">
            <div className="flex flex-col h-screen fixed bg-gray-950 text-white">
                <button className="m-4">
                    <BsHouseFill size='40px' className="m-auto"/>
                    <h1>Home</h1>
                </button>
                <button className="m-4">
                    <BsPersonFill size='40px' className="m-auto"/>
                    <h1>Clients</h1>
                </button>

            </div>
            <div className="grow h-screen bg-white">
                {
                    emails.map(data => {
                        return (
                            <div>
                                <h1>{data.title}</h1>
                                <p>{data.summary}</p>
                            </div>
                        )
                    })
                }
            </div>
        </div>
    </>
  )
}

export default Dashboard;