import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import CreateBlog from './components/createBlog'
import Home from './components/Home'
import Header from './components/Header'
import Landing from './components/landing'
import Register from './components/Register'

function App() {
  return (
    <div>
      <Router>
        <Header />
        <Routes>
          <Route exact path="/" element={<Landing />} />
          <Route path="/register" element={<Register />} />
          <Route path="/addblog" element={<CreateBlog />} />
          <Route path="/home" element={<Home />} />
        </Routes>
      </Router>
    </div>
  )
}

export default App
