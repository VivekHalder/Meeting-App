import { useEffect, useRef, useState } from 'react'

function Landing() {
  const [showMenu, setShowMenu] = useState(false)
  const menuRef = useRef<HTMLDivElement>(null)

  const toggleMenu = () => {
    setShowMenu(prev => !prev)
  }

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
        setShowMenu(false)
      }
    }

    document.addEventListener('mousedown', handleClickOutside)

    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
    }
  }, [])

  return (
   <div onClick={() => {
    if (showMenu) {
      toggleMenu()
    }
   }}>
    <input type="text" />
    <button>Enter</button>

    {/* New Meeting Button */}
    <div style={{ position: 'relative', display: 'inline-block' }}>
      <button onClick={toggleMenu}>New Meeting</button>

      {/* Dropdown Menu */}
      {
        showMenu && (
          <div
            ref={ menuRef }
            style={
              {
                position: 'absolute',
                top: '-100%',
                left: 0,
                backgroundColor: '#fff',
                boxShadow: '0px 4px 8px rgba(0, 0, 0, 0.1)',
                borderRadius: '8px',
                zIndex: 2,
                padding: '10px',
                width: '190%',
                color: 'black'
              }
            }
          >
            <div
              style={
                {
                  padding: '8px',
                  cursor: 'pointer'
                }
              }
            >
              Create a meeting for later
            </div>
            <div
              style={
                {
                  padding: '8px',
                  cursor: 'pointer'
                }
              }
            >
              Start an instant meeting
            </div>
            <div
              style={
                {
                  padding: '8px',
                  cursor: 'pointer'
                }
              }
            >
              Schedule in Google Calendar
            </div>
          </div>
        )
      }

    </div>
   </div> 
  )
}

export default Landing