import { signupAction } from './actions';
import Form from 'next/form';

export default function Signup() {
  return (
    <div className="flex flex-col items-center h-full w-[500px] p-4 mx-auto rounded-lg border border-black">
      <h1 className="text-lg font-bold">Signup</h1>
      <Form action={signupAction} className="grid grid-cols-[30%_70%] w-96 gap-2">


        <label htmlFor="username">Username:</label>
        <input className="rounded-md border border-black" type="text" id="username" name="username" required />



        <label htmlFor="password">Password:</label>
        <input className="rounded-md border border-black" type="password" id="password" name="password" required />



        <label htmlFor="fullname">Fullname:</label>
        <input className="rounded-md border border-black" type="text" id="fullname" name="fullname" required />



        <label htmlFor="email">Email:</label>
        <input className="rounded-md border border-black" type="email" id="email" name="email" required />



        <label htmlFor="phonenumber">Phone Number:</label>
        <input className="rounded-md border border-black" type="tel" id="phonenumber" name="phonenumber" required />



        <label htmlFor="gender">Gender:</label>
        <select className="rounded-md border border-black text-center" id="gender" name="gender" required>
          <option value="">Select Gender</option>
          <option value="male">Male</option>
          <option value="female">Female</option>
          <option value="other">Other</option>
        </select>

        <button type="submit" className="transition col-span-2 place-self-center w-36 h-8 rounded-md font-bold text-white bg-black hover:text-black hover:bg-white hover:ring-2 hover:ring-black duration-150">Login</button>
      </Form>
    </div>
  );
}