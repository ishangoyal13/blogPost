import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import CreateBlog from './components/createBlog'
import Home from './components/Home'

function App() {
  return (
    <div>
      <Router>
        <Routes>
          <Route path="/addblog" element={<CreateBlog />} />
          <Route path="/home" element={<Home />}/>
        </Routes>
      </Router>
    </div>
  )
}

export default App
