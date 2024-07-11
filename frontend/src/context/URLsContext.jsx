import { createContext,useReducer } from "react";

//create contexts
export const UrlContext = createContext();
export const UrlDispatchContext = createContext();

//initial state
const initialState = {
    urls: []
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
      default:
        return state;
    }
  };

// context provider
// eslint-disable-next-line react/prop-types
export function URLProvider({children}){
    const [state, dispatch] = useReducer(urlReducer, initialState);

    return (
      <UrlContext.Provider value={state}>
        <UrlDispatchContext.Provider value={dispatch}>
          {children}
        </UrlDispatchContext.Provider>
      </UrlContext.Provider>
    );

}
