import { render, Fragment, h, } from "fre"
import { useRoutes } from './use-route'
import './app.css'
import './m.css'
import Header from "./header/header"
import Footer from './header/footer'

const routes = {
    '/': import('./home/home'),
    '/login': import('./login/login'),
    '/register': import('./login/register'),
    '/publish/:id': import('./tm-upload/upload'),
    '/search/:k': import('./search/search'),
    '/watch/:pid':import('./book/book'),
    '/user/:id': import('./login/register'),
    '/recharge': import('./recharge/recharge'),
    '/my/:id': import('./admin/admin'),
    '/article/:id': import('./tm-article/article'),
    '/add-article/:pid': import('./tm-article/article'),
}

const App = () => {
    let route = useRoutes(routes)
    return <div>
        <Header />
        <div>{route}</div>
        {/* <Footer /> */}
    </div>
}

render(<App />, document.getElementById("app"))


