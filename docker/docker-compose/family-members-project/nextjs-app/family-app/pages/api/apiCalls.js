import { handleError, handleResponse } from "./apiUtils";

const memberUrl = process.env.NEXT_PUBLIC_API_URL + "/members";

export function getAllMembers() {
  return fetch(memberUrl, {
    method: "GET",
    headers: { "content-type": "application/json" },
  })
    .then(handleResponse)
    .catch(handleError);
}
