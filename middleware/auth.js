export default function ({ app, route, redirect }) {
  if (route.path !== '/') {
    if (!app.$fire.auth.currentUser) {
      return redirect('/')
    }
  } else if (route.path === '/') {
    if (app.$fire.auth.currentUser) {
      // leave them on the sign in page
    } else {
      return redirect('/')
    }
  }
}
