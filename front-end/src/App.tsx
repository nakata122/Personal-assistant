import { BrowserRouter, Route, Routes } from "react-router";
import Home from './views/Home';
import Dashboard from './views/Dashboard';
import Layout from "./shared/Layout";

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/" element={<Home />} />
          <Route path="/home" element={<Home />} />
          <Route path="/dashboard" element={<Dashboard />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
