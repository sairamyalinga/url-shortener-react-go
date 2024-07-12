import { useEffect,createContext,useReducer } from "react";
import API from "../lib/utils";

//create contexts
export const UrlContext = createContext();
export const UrlDispatchContext = createContext();

//initial state
const initialState = {
    urls: [],
    signin: !!localStorage.getItem('token')
  };

//define the reducer
const urlReducer = (state, action) => {
    switch (action.type) {
      case 'SET_URLS':
        return { ...state, urls: action.payload };
      case 'ADD_URL':
        return { ...state, urls: [...state.urls, action.payload] };
      case 'REMOVE_URL':
        return { ...state, urls: state.urls.filter(url => url.ShortID !== action.payload) };
      case 'SIGN_IN':
        return {...state, signin: action.payload};
      case 'LOG_OUT':
        return { ...state, signin: false, urls: [] }
      default:
        return state;
    }
  };

// context provider
// eslint-disable-next-line react/prop-types
export function URLProvider({children}){
    const [state, dispatch] = useReducer(urlReducer, initialState);
   

    const getURLs = () => {
        API.get("/urls")
        .then(response => {
          dispatch({ type: 'SET_URLS', payload: response.data.data });
        })
        .catch(error => {
          console.error("Error fetching URLs:", error);
        });

    }

    useEffect(() => {
        if (state.signin) {
          getURLs();
        }
      }, [state.signin, state.urls.length]);

    return (
      <UrlContext.Provider value={state}>
        <UrlDispatchContext.Provider value={dispatch}>
          {children}
        </UrlDispatchContext.Provider>
      </UrlContext.Provider>
    );

}
