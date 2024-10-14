import { createBrowserRouter, Navigate, RouterProvider } from "react-router-dom"
import Landing from "./components/Landing/Landing"

function App() {

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Navigate to="/landing" replace/>
    },
    {
      path: "/landing",
      element: <Landing/>
    }
  ])

  return (
    <>
      <RouterProvider router={router}/>
    </>
  )
}

export default App
