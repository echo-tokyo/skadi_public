import type { AppDispatch, RootState } from './store'

declare module 'react-redux' {
  type DefaultRootState = RootState

  export function useDispatch<TDispatch = AppDispatch>(): TDispatch
  export function useSelector<
    TState = RootState,
    TSelected = unknown,
  >(
    selector: (state: TState) => TSelected,
    equalityFn?: (left: TSelected, right: TSelected) => boolean,
  ): TSelected
}
