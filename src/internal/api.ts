export const API_URL = import.meta.env.VITE_API_URL;

export function fetchHomeResult() {
  return fetchJsonApi(`${API_URL}/home`);
}

export function fetchSearchResult(query: string) {
  return fetchJsonApi(`${API_URL}/search?q=${query}`);
}

export function fetchIdea(ideaId: number) {
  return fetchJsonApi(`${API_URL}/idea?idea=${ideaId}`);
}

export function publishIdea(data = {}) {
  return pushJsonApi(`${API_URL}/idea`, "POST", data);
}

export function patchIdea(ideaId: number, data = {}) {
  return pushJsonApi(`${API_URL}/idea?idea=${ideaId}`, "PATCH", data);
}

export function deleteIdea(ideaId: number) {
  return pushJsonApi(`${API_URL}/idea?idea=${ideaId}`, "DELETE");
}

export function patchVote(ideaId: number, voteValue: number) {
  return pushJsonApi(`${API_URL}/vote?idea=${ideaId}`, "PATCH", {
    vote: voteValue,
  });
}

export function fetchJsonApi(url: string) {
  return new Promise((res, rej) => {
    fetch(url, {
      method: "GET",
      headers: {
        Accept: "application/json, text/plain, */*",
      },
    })
      .then((resp) => {
        resp
          .json()
          .then((r) => {
            res(r);
          })
          .catch(function (err) {
            rej(err);
          });
      })
      .catch(() => {
        rej(null);
      });
  });
}

export function pushJsonApi(
  url: string,
  method: "POST" | "PATCH" | "DELETE",
  data = {}
) {
  return new Promise((res, rej) => {
    fetch(url, {
      method: method,
      headers: {
        Accept: "application/json, text/plain, */*",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((resp) => {
        resp
          .json()
          .then((r) => {
            res(r);
          })
          .catch(function (err) {
            rej(err);
          });
      })
      .catch(() => {
        rej(null);
      });
  });
}
