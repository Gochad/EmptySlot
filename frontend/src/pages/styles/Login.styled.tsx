import styled from 'styled-components';

export const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #aec2dc;

  h2 {
    text-align: center;
    margin-bottom: 20px;
    color: #333;
  }
`;

export const Form = styled.form`
  padding: 20px;
  max-width: 400px;
  width: 100%;
  background-color: white;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;

  label {
    display: block;
    margin-bottom: 5px;
    color: #666;
  }

  input {
    width: 100%;
    padding: 10px 0;
    margin-bottom: 20px;
    border: 1px solid #ccc;
    border-radius: 4px;
  }

  p.error {
    color: red;
    text-align: center;
  }
`;

export const Button = styled.button`
  width: 100%;
  padding: 10px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  font-size: 16px;
  cursor: pointer;

  :hover {
    background-color: #0056b3;
  }
`;
