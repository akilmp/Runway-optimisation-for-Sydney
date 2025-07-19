import { configureStore, createSlice } from '@reduxjs/toolkit';

interface SeatState {
  passengers: string[];
}

const initialState: SeatState = {
  passengers: [],
};

const seatSlice = createSlice({
  name: 'seats',
  initialState,
  reducers: {
    addPassenger: (state, action) => {
      state.passengers.push(action.payload);
    },
  },
});

export const { addPassenger } = seatSlice.actions;

export const store = configureStore({
  reducer: {
    seats: seatSlice.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
