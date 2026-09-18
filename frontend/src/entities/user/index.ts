export {
  setUserData,
  clearUserData,
  selectAuthenticatedUser,
} from './model/user-slice.ts'

export type { IUserResponse } from './model/types.ts'

export { default as userReducer } from './model/user-slice.ts'

export { useGetMeQuery } from './api/user-api.ts'
