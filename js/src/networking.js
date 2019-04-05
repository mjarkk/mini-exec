let key = ''

const f = async (what, data) => {
  const req = await fetch(location.href, {
    method: 'POST',
    body: JSON.stringify(Object.assign({key, what}, data)),
    headers: {
      'Content-Type': 'application/json' 
    }
  })
  const res = await req.json()
  return res 
}

const actions = {
  logout() {
    key = ''
  },
  login(newKey) {
    key = newKey
    return f('init')
  },
  update() {
    return f('update')
  }
}

export default actions
