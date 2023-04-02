import { createStore } from 'vuex'

const store = createStore({
    state () {
        return {
            user: null,
            token: null,
        }
    },
    mutations: {
        setUser(state, user) {
            state.user = user;
        },
        setToken(state, token) {
            state.token = token;
        },
    },
    getters: {
        isLoggedIn(state) {
            return !!state.token;
        },
        token(state) {
            return state.token;
        },
    }
});

export default store;