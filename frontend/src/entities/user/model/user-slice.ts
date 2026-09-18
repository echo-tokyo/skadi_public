import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IUser, IUserResponse } from './types'

interface IUserState {
  user: IUser | null
  isAuthenticated: boolean
  authChecked: boolean
}

const initialState: IUserState = {
  user: null,
  isAuthenticated: false,
  authChecked: false,
}

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUserData: (state, { payload }: PayloadAction<IUserResponse>) => {
      state.user = payload
      state.isAuthenticated = true
      state.authChecked = true
    },
    clearUserData: (state) => {
      state.user = null
      state.isAuthenticated = false
      state.authChecked = true
    },
  },
})

export const { setUserData, clearUserData } = userSlice.actions
export default userSlice.reducer
export const selectAuthenticatedUser = (state: { user: IUserState }): IUser =>
  state.user.user!
