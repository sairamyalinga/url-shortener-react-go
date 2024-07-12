import { useContext } from "react"
import PropTypes from "prop-types";
import Header from "./Header";
import API, { handleCopyToClipboard } from "../lib/utils";
import {UrlContext, UrlDispatchContext} from "../context/URLsContext"

function Urlitem({ url, handleDelete }) {
  const shorturl = `${import.meta.env.VITE_API_URL}/${url.ShortID}`;
  
  return (
    <div className="card border border-0 mb-3 mx-5">
      <div className="card-body rounded border border-info bg-info-subtle text-info-emphasis">
        <div className="d-flex justify-content-between align-items-center  ">
          <p className="mb-0">
            Short URL: <a href = {shorturl} target="_blank">{shorturl}</a>
          </p>
          <div>
            <button
              className="btn btn-secondary btn-sm mr-2 mx-4"
              onClick={() => handleCopyToClipboard(shorturl)}
            >
              Copy
            </button>
            <button
              className="btn btn-danger btn-sm"
              onClick={() => handleDelete(url.ShortID)}
            >
              Delete
            </button>
          </div>
        </div>
        <p className="mb-0">
          Original URL: <a href={`${url.Url}`}>{url.URL}</a>
        </p>
      </div>
    </div>
  );
}

function Urlboard() {
  const { urls } = useContext(UrlContext);
  const dispatch = useContext(UrlDispatchContext);
  

  const handleDelete = (shortID) => {
    
    API
      .delete("/url", {
        data:{
          shortID,
        }
      })
      .then(() => {
        dispatch({ type: 'REMOVE_URL', payload: shortID });
      })
      .catch((error) => {
        alert(error);
      });
  };
  
  return (
    <div>
      <Header page="urlboard" />
      {urls.length > 0 ? (
        urls.map((url, index) => <Urlitem key={index} url={url} handleDelete={() => handleDelete(url.ShortID)} />)
      ) : (
        <div className="d-flex align-items-center justify-content-center">
          <div className="badge text-bg-info text-wrap fs-5 fst-italic">
            {" "}
            No shortURLs in your account.
          </div>
        </div>
      )}
    </div>
  );
}


Urlitem.propTypes = {
  url: PropTypes.object.isRequired,
  handleDelete: PropTypes.func.isRequired
};

export default Urlboard;
