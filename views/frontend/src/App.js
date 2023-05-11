import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import CreateBlog from './components/createBlog'
import Home from './components/Home'
import Register from './components/Register'
import Login from './components/Login'

function App() {
  return (
    <div>
      <Router>
        <Routes>
          <Route exact path="/" element={<Register />} />
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/addblog" element={<CreateBlog />} />
          <Route path="/home" element={<Home />} />
        </Routes>
      </Router>
    </div>
  )
}

export default App
