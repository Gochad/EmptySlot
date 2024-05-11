import styled from "styled-components";
import Modal from "react-modal";

export const Form = styled.form`
  padding: 20px;
  max-width: 400px;
  width: 100%;
  background-color: #fff;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border: 1px solid #e0e0e0;
  border-radius: 10px;
  font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;

  label {
    display: block;
    margin-bottom: 8px;
    color: #333;
    font-size: 14px;
  }

  input {
    width: 100%;
    padding: 12px 10px;
    margin-bottom: 15px;
    border: 1px solid #ccc;
    border-radius: 6px;
    font-size: 16px;
    transition: border-color 0.3s;

    &:focus {
      border-color: #0056b3;
      outline: none;
    }
  }

  p.error {
    color: #ff0000;
    text-align: center;
    margin-top: -10px;
    font-size: 14px;
  }
`;

export const ModalStyled = styled(Modal)`
  min-height: 10vh;
  min-width: 400px;
  max-height: 90vh;
  max-width: 90%;
  margin: auto;
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;
  border-radius: 8px;
  border: 1px solid #ccc;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
  overflow: auto;
  transition: all 0.3s ease-out;

  &:focus {
    outline: none;
  }

  .Modal__overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 50;
    transition: opacity 0.3s ease-in-out;
  }
`;
