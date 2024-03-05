import styled from "styled-components";

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

export const modalStyles = {
    content: {
        top: '50%',
        left: '50%',
        right: 'auto',
        bottom: 'auto',
        marginRight: '-50%',
        transform: 'translate(-50%, -50%)',
        backgroundColor: 'papayawhip',
    },
};