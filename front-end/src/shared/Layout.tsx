import Header from './Header';
import { Outlet } from 'react-router';

function Layout() {

  return (
    <>
        <Header />
        <main className='pt-15'>
            <Outlet />
        </main>
    </>
  )
}

export default Layout;