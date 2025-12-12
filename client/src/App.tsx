import { BrowserRouter, Route, Routes } from "react-router";
import Home from './views/Home';
import Dashboard from './views/dashboard/Dashboard';
import Layout from "./shared/Layout";

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/" element={<Home />} />
          <Route path="/home" element={<Home />} />
        </Route>
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
